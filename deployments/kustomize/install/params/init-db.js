const mongoHost = process.env.BLOOD_DONORS_API_MONGODB_HOST
const mongoPort = process.env.BLOOD_DONORS_API_MONGODB_PORT

const mongoUser = process.env.BLOOD_DONORS_API_MONGODB_USERNAME
const mongoPassword = process.env.BLOOD_DONORS_API_MONGODB_PASSWORD

const database = process.env.BLOOD_DONORS_API_MONGODB_DATABASE
const collection = process.env.BLOOD_DONORS_API_MONGODB_COLLECTION

const retrySeconds = parseInt(process.env.RETRY_CONNECTION_SECONDS || "5") || 5;

// try to connect to mongoDB until it is not available
let connection;
while(true) {
    try {
        connection = Mongo(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`);
        break;
    } catch (exception) {
        print(`Cannot connect to mongoDB: ${exception}`);
        print(`Will retry after ${retrySeconds} seconds`)
        sleep(retrySeconds * 1000);
    }
}

// if database and collection exists, exit with success - already initialized
const databases = connection.getDBNames()
if (databases.includes(database)) {
    const dbInstance = connection.getDB(database)
    collections = dbInstance.getCollectionNames()
    if (collections.includes(collection)) {
      print(`Collection '${collection}' already exists in database '${database}'`)
        process.exit(0);
    }
}

// initialize
// create database and collection
const db = connection.getDB(database)
db.createCollection(collection)

// create indexes
db[collection].createIndex({ "id": 1 })

//insert sample data
let result = db[collection].insertMany([
    {
        "id": "bratislava-bory",
        "name": "Bratislava Bory",
        "address": "Cesta mládeže 2/A, Bratislava",
        "predefinedDonationTypes": [
            { "value": "Darovanie krvi", "code": "blood", "typicalDurationMinutes": 15 },
            { "value": "Darovanie krvnej plazmy", "code": "plasma", "typicalDurationMinutes": 45 }
        ],
        "donors": [
            {
                "id": "x321ab3",
                "name": "Adrián Vančo",
                "donorId": "1234567890",
                "sex": "M",
                "bloodType": "A+",
                "email": "vancoadrian7@gmail.com",
                "phone": "+421917529741",
                "preferredDonationType": "blood",
                "preferredSite": "bratislava-bory",
                "eligible": true,
                "registeredSince": new Date("2026-05-20T09:00:00Z"),
                "donations": [
                    { "date": new Date("2025-07-10T09:00:00Z"), "donationType": { "value": "Darovanie krvi", "code": "blood" }, "status": "Odber dokončený" },
                    { "date": new Date("2024-02-13T10:45:00Z"), "donationType": { "value": "Darovanie krvi", "code": "blood" }, "status": "Odber dokončený" }
                ]
            }
        ]
    },
    {
        "id": "malacky",
        "name": "Malacky",
        "address": "Allianz/Galéria skla, Cesta mládeže 2/A, Malacky",
        "predefinedDonationTypes": [
            { "value": "Darovanie krvi", "code": "blood", "typicalDurationMinutes": 15 },
            { "value": "Darovanie krvnej plazmy", "code": "plasma", "typicalDurationMinutes": 45 }
        ],
        "donors": []
    }
]);

if (result.writeError) {
    console.error(result)
    print(`Error when writing the data: ${result.errmsg}`)
}

// exit with success
process.exit(0);