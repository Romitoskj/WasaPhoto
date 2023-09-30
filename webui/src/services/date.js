function displayDateAndTime(date) {
	return date.toDateString() + " at " + date.toLocaleTimeString('it-IT');
}

export default displayDateAndTime;
