<!DOCTYPE html>
<html>
<head>
<style>
body {font-family: Arial, Helvetica, sans-serif;margin-top: 150px;}
input[type=text], input[type=password] {
  width: 100%;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  box-sizing: border-box;
}
button {
  background-color: #e9711c;
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  cursor: pointer;
  width: 100%;
  font-size: 18px;
  font-weight: 800;
}
button:hover {
  opacity: 0.8;
}
.container {
  padding: 16px;
}
h1{
  margin: auto;
  width: 50%;
  text-align: center;
}
</style>
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
