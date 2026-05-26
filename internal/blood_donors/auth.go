package blood_donors

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Rola pracovníka transfúznej stanice. Hodnota musí zodpovedať role priradenej
// v OPA politike (blood-donors-gitops/infrastructure/opa-plugin/params/policy.rego).
const roleWorker = "pracovnik"

// Zoznam emailov pracovníkov - musí ostať zhodný s worker_emails v policy.rego
// a WORKER_EMAILS v blood-donors-ufe/src/global/auth.ts.
// Slúži ako spoľahlivý fallback: na zdieľanom klastri nemusí spoločná OPA
// posielať aplikačnú rolu "pracovnik", preto pracovníka rozpoznáme aj z emailu.
var workerEmails = map[string]bool{
	"xvancoa@stuba.sk":  true,
	"qunger@stuba.sk":   true,
	"qmicuch@stuba.sk":  true,
	"qsevcikm@stuba.sk": true,
	"qhudakm1@stuba.sk": true,
}

// emailIsWorker zistí, či je email prihláseného používateľa v zozname pracovníkov.
func emailIsWorker(c *gin.Context) bool {
	email := strings.ToLower(forwardedEmail(c))
	return email != "" && workerEmails[email]
}

// isWorker - pracovník podľa role z OPA alebo podľa emailu (fallback).
func isWorker(c *gin.Context) bool {
	return hasRole(c, roleWorker) || emailIsWorker(c)
}

// forwardedRoles vráti role, ktoré požiadavke priradila OPA cez gateway
// (hlavička "x-forwarded-roles", napr. "darca" alebo "pracovnik, monitoring").
func forwardedRoles(c *gin.Context) []string {
	raw := c.GetHeader("x-forwarded-roles")
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	roles := make([]string, 0, len(parts))
	for _, p := range parts {
		if r := strings.TrimSpace(p); r != "" {
			roles = append(roles, r)
		}
	}
	return roles
}

// hasRole zistí, či používateľ má danú rolu.
func hasRole(c *gin.Context, role string) bool {
	for _, r := range forwardedRoles(c) {
		if r == role {
			return true
		}
	}
	return false
}

// forwardedEmail vráti email prihláseného používateľa (od gateway).
func forwardedEmail(c *gin.Context) string {
	return strings.TrimSpace(c.GetHeader("x-forwarded-email"))
}

// behindGateway je true, ak požiadavka prešla cez gateway (má identitu z OIDC -
// hlavičku s emailom alebo rolami). Ak chýbajú obe, ide o lokálny dev / priamy
// prístup mimo gateway a autorizáciu nevynucujeme - v produkcii je webapi
// dostupné výhradne cez gateway, ktorý neautentifikované požiadavky odmietne.
func behindGateway(c *gin.Context) bool {
	return forwardedEmail(c) != "" || c.GetHeader("x-forwarded-roles") != ""
}

// requireWorker zamietne požiadavku so stavom 403, ak používateľ nie je
// pracovník. Vráti true, ak smie pokračovať.
func requireWorker(c *gin.Context) bool {
	if !behindGateway(c) || isWorker(c) {
		return true
	}
	c.JSON(http.StatusForbidden, gin.H{
		"status":  http.StatusForbidden,
		"message": "Akcia vyžaduje rolu pracovníka",
	})
	return false
}
