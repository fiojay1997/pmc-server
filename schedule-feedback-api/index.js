const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const db = require('./queries');
const PORT = process.env.PORT || 5000;

// parse the body into json
app.use(bodyParser.json());
app.use(
  bodyParser.urlencoded({
    extended: true,
  })
);

app.listen(PORT, () => {
  console.log(`App running on port ${PORT}.`)
});

app.get('/', (request, response) => {
  response.json({ info: 'schedule API' })
});

// schedule API
app.get('/schedule/:user_id/:semester_id', db.getSchedule);
app.post('/schedule/add', db.addToSchedule);
app.post('/schedule/remove', db.removeFromSchedule);

// feedback API
app.get('/feedback', db.getFeedbacks);
app.get('/feedback/:id', db.getFeedbackById);
app.post('/feedback', db.createFeedback);

