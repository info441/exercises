# Web Client for helloworld

Ensure that your web API server in the `../helloserver/` directory is running and ready for requests. Then create a new web
 app in this directory that meets the following requirements:

- Your page must have an `<input>` element into which the user can type their name. 
- Your page must have a `<button>` element the user can click/tap to submit that data. You may also let the user
 trigger a submit by hitting the `Enter` key while focus is in the `<input>` element (hint: use a `<form>` element and
  catch the form's `submit` event).
- When the user submits the data, do two things:
	- Use the [Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch) to make a 
	`GET /?name={name-entered}` request to your web 	server, where `{name-entered}` is replaced by the value of your `<input>` element. Render the text you get back into some element on 	the page so that the user can see the 
	welcome message. 
	- Use an `<img>` element to display the identicon for the name in your `<input>` element. Remember that the 
	`/identicon/{name-entered}` resource path of your web server will return the identicon PNG image bytes for the name
	 you provide. Therefore, you can just set the response of `/identicon/{name-entered}` to the `<img>` element's source.
	
## Running Local Testing Server
I strongly encourage you to install `live-server` as you will be needing a development server throughout the course.
`live-server` has live reloading capability, therefore it'll save you time and accelerate your development.

### Installation 
You need node.js and npm. You should probably install this globally. Run the command below in your terminal.
```
npm install -g live-server
```
### Usage 
Navigate to your project's directory. Run the command below in your terminal.
```
live-server
```
### Documentation 
Further documentation on `live-server` can be found [here](https://www.npmjs.com/package/live-server)