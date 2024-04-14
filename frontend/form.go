package frontend

// HTML template for the form
var FormTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Confirmation Form</title>
    <link rel="stylesheet" type="text/css" href="/static/style.css">
</head>
<body>
    <h1>Confirmation Form</h1>
    <form action="/confirm" method="post">
        <input type="hidden" id="uuid" name="uuid">
		<label for="firstname">Nome:</label>
		<input type="text" id="firstname" name="firstname" required>
        <br>
		<label for="lastname">Cognome:</label>
		<input type="lastname" id="guests" name="lastname" required>
        <br>
        <label for="guests">Numero <b>totale</b> di partecipanti:</label>
        <input type="number" id="guests" name="guests" required>
        <br>
        <input type="submit" value="Confirm">
    </form>	
	<script>
	// Function to extract query parameter from URL
	function getQueryParam(name) {
		const urlParams = new URLSearchParams(window.location.search);
		return urlParams.get(name);
	}

	// Set the UUID value from query string to the hidden input field
	const uuidInput = document.getElementById('uuid');
	const uuidValue = getQueryParam('uuid');
	uuidInput.value = uuidValue;

	// Submit the form
	document.getElementById('confirmationForm').onsubmit = function() {
		if (!uuidValue) {
			alert('UUID not found in query string');
			return false; // Prevent form submission if UUID is missing
		}
	};
</script>
</body>
</html>
`
