// ./src/index.js

// import dependencies
const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const helmet = require('helmet');
const morgan = require('morgan');

// import database
const {startDatabase} = require('./database/database');
const {insertBud, getBuds, deleteBud, updateBud} = require('./database/buds');

// define the express app
const app = express();

// adding helmet to enhance API's security
app.use(helmet());

// use bodyParser to parse JSON bodies
app.use(bodyParser.json());

// enabling CORS (Cross-Origin Resource Sharing)
// for all requests
app.use(cors());

// add morgan to log HTTP requests
app.use(morgan('combined'));

// define endpoints
app.get('/', async(req, res) => {
  res.send(await getBuds());
});

app.post('/', async(req, res) => {
    const newBud = req.body;
    await insertBud(newBud);
    res.send({message: "new strain has been added to stash"});
});

app.delete('/:id', async(req, res) => {
    await deleteBud(req.params.id);
    res.send({message: "strain has been depleted; removed from stash :("});
});

app.put('/:id', async(req, res) => {
    const updatedBud = req.body;
    await updateBud(req.params.id, updatedBud);
    res.send({message: "updated strain in stash"});
});

startDatabase().then(async () => {
    await insertBud({title: "Platfinum Kush", price: 12});

    // starting the server
    app.listen(3001, () => {
	console.log("listening to server on port 3001");
    });
});
