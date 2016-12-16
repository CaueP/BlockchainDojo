/*eslint-env node*/

//------------------------------------------------------------------------------
// node.js starter application for Bluemix
//------------------------------------------------------------------------------

// This application uses express as its web server
// for more info, see: http://expressjs.com
var express = require('express');
var bodyParser = require('body-parser')


// cfenv provides access to your Cloud Foundry environment
// for more info, see: https://www.npmjs.com/package/cfenv
var cfenv = require('cfenv');

// create a new express server
var app = express();

// serve the files out of ./public as our main files
app.use(express.static(__dirname + '/public'));

app.use(bodyParser.json());
//Declarando um logger para verificar as rotas que estao sendo requisitadas.
const logger = require('morgan'); 
app.use(logger('dev')); 


// get the app environment from Cloud Foundry
var appEnv = cfenv.getAppEnv();


/// rotas

app.get('/', function (req, res) {
  res.send('Hello World!');
}); 

app.post('/atualizar', function (req, res) {
		console.log("/atualizar: " + JSON.stringify(req.body  ));
    res.send(JSON.stringify(req.body));
	})

// start server on the specified port and binding host
app.listen(appEnv.port, '0.0.0.0', function() {
  // print a message when the server starts listening
  console.log("server starting on " + appEnv.url);
});
