# Web Client for zip
Ensure that your web API server in the `../zipserver/` directory is running and ready for requests. 
then create a new web app in this directory that meets the following requirements:
- Your application must have an `<input>` element where the user can enter a city
- Your application must have **two** `<select>` elements that get auto-populated based on the results
returned from your server.
- Only send a request when the user is finished typing a city, **don't** send a request on every keystroke. 

A few tips:
- Use the `"change"` event on the city input to catch when the user is done entering the city name.
 This event is triggered when the input looses focus. You can use `"input"` to respond to each keystroke,
  but that will generate a lot of requests that simply return null results. See https://schier.co/blog/2014/12/08/wait-for-user-to-stop-typing-using-javascript.html
  for an example.
- Remember that your server will respond with a JSON array of objects, each of which will have `code`, `city` and 
`state` properties.
- The same `state` may appear multiple times in the results, so you will need to populate the state `<select>` with 
just the distinct state names.
- Beware that if the city is not found, your server will return a JSON `null`, not an empty array.
