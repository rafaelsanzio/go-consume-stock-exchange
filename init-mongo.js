db.createUser({
  user: "root",
  pwd: "root",
  roles: [
    { role: "readAnyDatabase", db: "stockexchange" },
    { role: "dbAdminAnyDatabase", db: "stockexchange" },
    { role: "userAdminAnyDatabase", db: "stockexchange" },
  ],
});

let db = connect("localhost:27017/stockexchange");
db.auth("root", "root");

db.createCollection("events");
