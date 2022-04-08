const express = require('express')
const cors = require('cors')
const app = express()

// Middleware
app.use(cors())
app.use(express.json())

// Routes
app.use('/', require('./routes/comments'))


app.listen(9000, ()=> console.log('Server listening on port 9000'))