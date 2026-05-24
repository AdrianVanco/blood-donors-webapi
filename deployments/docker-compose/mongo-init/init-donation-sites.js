// Seed pre lokálne testovanie - spustí sa automaticky pri prvom štarte Monga
// (mongo image vykoná skripty z /docker-entrypoint-initdb.d na prázdnom dátovom priečinku).
db = db.getSiblingDB('cv2xvancoa-blood-donors');
db.createCollection('donation_sites');
db.donation_sites.createIndex({ id: 1 });

db.donation_sites.insertMany([
  {
    id: "bratislava-bory",
    name: "Bratislava Bory",
    address: "Cesta mládeže 2/A, Bratislava",
    predefinedDonationTypes: [
      { value: "Darovanie krvi", code: "blood", typicalDurationMinutes: 15 },
      { value: "Darovanie krvnej plazmy", code: "plasma", typicalDurationMinutes: 45 }
    ],
    donors: [
      {
        id: "x321ab3",
        name: "Adrián Vančo",
        donorId: "1234567890",
        sex: "M",
        bloodType: "A+",
        email: "vancoadrian7@gmail.com",
        phone: "+421917529741",
        preferredDonationType: "blood",
        preferredSite: "bratislava-bory",
        eligible: true,
        registeredSince: ISODate("2026-05-20T09:00:00Z"),
        donations: [
          { date: ISODate("2025-07-10T09:00:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Odber dokončený" },
          { date: ISODate("2024-02-13T10:45:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Odber dokončený" }
        ]
      },
      {
        id: "y998kl2",
        name: "Jana Nováková",
        donorId: "9876543210",
        sex: "F",
        bloodType: "0-",
        email: "jana.novakova@example.com",
        phone: "+421900111222",
        preferredDonationType: "plasma",
        preferredSite: "malacky",
        eligible: false,
        eligibilityNote: "Chronické ochorenie, darovanie nie je možné",
        registeredSince: ISODate("2026-05-21T10:00:00Z"),
        donations: []
      }
    ]
  },
  {
    id: "malacky",
    name: "Malacky",
    address: "Allianz/Galéria skla, Cesta mládeže 2/A, Malacky",
    predefinedDonationTypes: [
      { value: "Darovanie krvi", code: "blood", typicalDurationMinutes: 15 },
      { value: "Darovanie krvnej plazmy", code: "plasma", typicalDurationMinutes: 45 }
    ],
    donors: []
  }
]);
