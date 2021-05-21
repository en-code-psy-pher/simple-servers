// ./src/database/ads.js

const {getDatabase} = require('./database');
const {ObjectID} = require('mongodb');

const collectionName = "buds";

async function insertBud(bud) {
    const database = await getDatabase();
    const {insertID} = await database.collection(collectionName).insertOne(bud);
    return insertID;
}

async function getBuds() {
    const database = await getDatabase();
    return await database.collection(collectionName).find({}).toArray();
}

async function deleteBud(id) {
    const database = await getDatabase();
    await database.collection(collectionName).deleteOne({
	_id: new ObjectID(id),
    });
}

async function updateBud(id, bud) {
    const database = await getDatabase();
    delete bud._id;
    await database.collection(collectionName).update(
	{_id: new ObjectID(id), },
	{
	    $set: {
		...bud,
	    },
	},
    );
}


module.exports = {
    insertBud,
    getBuds,
    deleteBud,
    updateBud,
};
