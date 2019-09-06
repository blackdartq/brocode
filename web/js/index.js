function myFuction(){
	var para = document.querySelector("p");
	var navbar = document.querySelector("navbar");
	para.innerHTML = "Test";
	for(var x=0; x<5; x++){
		navbar.write("<button>IDK i hope this works</button>");
		//document.write("<h2>fuck the police</h2>");
	}
}

function countDown(){
	var minutesInput = document.getElementById("minutes-input");
	var minutesDiv = document.getElementById("minutes-div");
	minutesDiv.innerHTML = minutesInput.innerHTML
	var secondsDiv = document.getElementById("seconds-div");
	/*for(var min = 50; min > 0; min--){
		minutesDiv.innerHTML = min;
		console.log(min);
		for(var seconds = 1; seconds < 60; seconds++){
			secondsDiv.innerHTML = seconds;
		}
	}
	*/

}

function logContactForm(){
    var contactForm = {
        name: document.querySelector("input[name='name']").value,
        email: document.querySelector("input[name='email']").value,
        message: document.querySelector("input[name='message']").value
    }

    alert("your information is: " + contactForm.name + " \n email: " + contactForm.email + " \n message: " + contactForm.message)
    console.log("your information is: " + contactForm.name + " \n email: " + contactForm.email + " \n message: " + contactForm.message)
}


