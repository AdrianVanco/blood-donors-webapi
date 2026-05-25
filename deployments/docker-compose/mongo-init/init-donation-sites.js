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
        registeredSince: ISODate("2024-02-01T09:00:00Z"),
        donations: [
          { date: ISODate("2025-07-10T09:00:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Odber dokončený" },
          { date: ISODate("2024-02-13T10:45:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Odber dokončený" },
          { date: ISODate("2026-06-18T09:00:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Rezervácia dokončená" },
          { date: ISODate("2026-03-02T09:00:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Zrušená rezervácia", note: "Darca sa nedostavil" }
        ]
      },
      {
        id: "pethor01",
        name: "Peter Horváth",
        donorId: "1112223334",
        sex: "M",
        bloodType: "B+",
        email: "peter.horvath@example.com",
        phone: "+421901234567",
        preferredDonationType: "blood",
        preferredSite: "bratislava-bory",
        eligible: true,
        registeredSince: ISODate("2025-01-15T09:00:00Z"),
        donations: [
          { date: ISODate("2026-04-20T09:15:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Odber dokončený" }
        ]
      },
      {
        id: "markov02",
        name: "Mária Kováčová",
        donorId: "2223334445",
        sex: "F",
        bloodType: "A-",
        email: "maria.kovacova@example.com",
        phone: "+421902345678",
        preferredDonationType: "blood",
        preferredSite: "bratislava-bory",
        eligible: true,
        registeredSince: ISODate("2024-06-01T09:00:00Z"),
        donations: [
          { date: ISODate("2026-01-08T09:00:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Odber dokončený" },
          { date: ISODate("2026-03-20T09:00:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Odber dokončený" },
          { date: ISODate("2026-05-19T09:00:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Odber dokončený" }
        ]
      },
      {
        id: "luktot03",
        name: "Lukáš Tóth",
        donorId: "3334445556",
        sex: "M",
        bloodType: "AB+",
        email: "lukas.toth@example.com",
        phone: "+421903456789",
        preferredDonationType: "plasma",
        preferredSite: "bratislava-bory",
        eligible: true,
        registeredSince: ISODate("2025-09-10T09:00:00Z"),
        donations: [
          { date: ISODate("2026-05-18T11:00:00Z"), donationType: { value: "Darovanie krvnej plazmy", code: "plasma" }, status: "Odber dokončený" }
        ]
      },
      {
        id: "evabal04",
        name: "Eva Balážová",
        donorId: "4445556667",
        sex: "F",
        bloodType: "B-",
        email: "eva.balazova@example.com",
        phone: "+421904567890",
        preferredDonationType: "both",
        preferredSite: "bratislava-bory",
        eligible: true,
        registeredSince: ISODate("2026-05-10T09:00:00Z"),
        donations: []
      },
      {
        id: "jannov05",
        name: "Jana Nováková",
        donorId: "9876543210",
        sex: "F",
        bloodType: "0-",
        email: "jana.novakova@example.com",
        phone: "+421900111222",
        preferredDonationType: "plasma",
        preferredSite: "bratislava-bory",
        eligible: false,
        eligibilityNote: "Chronické ochorenie, darovanie nie je možné",
        registeredSince: ISODate("2025-11-02T09:00:00Z"),
        donations: [
          { date: ISODate("2025-12-05T09:00:00Z"), donationType: { value: "Darovanie krvi", code: "blood" }, status: "Nespôsobilý", note: "Nízky hemoglobín" }
        ]
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
