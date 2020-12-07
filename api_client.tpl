<!DOCTYPE html>
<html>
<head>
</head>
<body>
	<h1>Safe Harbor De-Identification</h1>
	<form action="/deidentify" method="post">
		<div class="container">
			<label for="birthdate"><b>BirthDate</b></label>
			<input type="text" placeholder="Enter BirthDate" name="birthdate" required>

			<label for="psw"><b>ZipCode</b></label>
			<input type="text" placeholder="Enter Zip code" name="zipcode" required>

			<label for="admissiondate"><b>Admission Date</b></label>
			<input type="text" placeholder="Enter Admission Date" name="admissiondate" required>

			<label for="dischargedate"><b>Discharge Date</b></label>
			<input type="text" placeholder="Enter Discharge Date" name="dischargedate" required>

			<label for="notes"><b>Notes</b></label>
			<input type="text" placeholder="Notes..." name="notes">

			<button type="submit">Execute De-identification</button>
		</div>
	</form>
</body>
</html>
