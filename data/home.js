
const cards = document.querySelectorAll(".card__inner");
for (let i = 0; i < cards.length; i++) {
    const card = cards[i];
    card.addEventListener("click", () =>
    {
        card.classList.toggle("is-flipped");

    });
}

function sendForm() {
    const XHR = new XMLHttpRequest();
    const FD = new FormData();
    var select = document.getElementById('filterByColor');
    var value = select.options[select.selectedIndex].value;
    member = document.getElementsByName('members');
    var checkBox = "";
    for (var i = 0; i < member.length; i++) {
        if (member[i].checked) {
            checkBox += member[i].value
        }
    }
    var startYear = document.getElementById("inputRange").value

    FD.append("input", document.getElementById("myInput").value)
    FD.append("country", value)
    FD.append("nbMembers", checkBox)
    FD.append("startYear", startYear)

    XHR.open('POST',"/")
    XHR.send(FD)

    location.assign("/")
}
function myFunction() {
    // Declare variables
    var input, filter, ul, li, a, i, txtValue;
    input = document.getElementById('myInput');
    filter = input.value.toUpperCase();
    ul = document.getElementById("myUL");
    li = ul.getElementsByTagName('li');
    console.log(filter)
    // Loop through all list items, and hide those who don't match the search query
    for (i = 0; i < li.length; i++) {
        a = li[i].getElementsByTagName("a")[0];
        txtValue = a.textContent || a.innerText;

        if (txtValue.toUpperCase().indexOf(filter) > -1) {
            li[i].style.display = "";
        } else {
            li[i].style.display = "none";
            }
    }
}

var checkList = document.getElementById('list1');
checkList.getElementsByClassName('anchor')[0].onclick = function() {
    if (checkList.classList.contains('visible'))
        checkList.classList.remove('visible');
    else
        checkList.classList.add('visible');
}
var check = document.getElementById('carrerStart');
check.getElementsByClassName('carrer')[0].onclick = function() {

    if (check.classList.contains('visibility'))
        check.classList.remove('visibility');
    else
        check.classList.add('visibility');
}

//
const inputField = document.getElementById('myInput');


    // Load the stored value from localStorage when the page loads
    document.addEventListener('DOMContentLoaded', () => {
        const storedValue = localStorage.getItem('inputValue');
        if (storedValue) {
            inputField.value = storedValue;
        }
    });

// Save the current value to localStorage when the input value changes
    inputField.addEventListener('input', () => {
        localStorage.setItem('inputValue', inputField.value);
    });

var closeMapButton = document.getElementById('close-map-button');
closeMapButton.addEventListener('click', hideMap);

function hideMap() {
    var mapContainer = document.querySelector("#map");
    mapContainer.style.opacity = "0";
    mapContainer.style.visibility = "hidden";
    mapContainer.style.transform = "translate(-50%, -50%) scale(0)"; // Modifiez cette ligne

    var btnContainer = document.querySelector("#close-map-button");
    btnContainer.style.opacity = "0";
    btnContainer.style.visibility = "hidden";
    btnContainer.style.transform = "translate(-50%, -50%) scale(0)";
}


var map;
var markers = [];

function initMap() {
    var options = {
        zoom: 2,
        center: {lat: 0, lng: 0},
    };
    map = new google.maps.Map(document.getElementById("map"), options);
}

function showMarkers(id) {

    for (var i = 0; i < markers.length; i++) {
        markers[i].setMap(null);
    }
    markers = [];


    var locationDiv = document.querySelector('[data-id="' + id + '"]');
    var locationPars = locationDiv.getElementsByTagName('p');
    for (var i = 0; i < locationPars.length; i++) {
        var locationText = locationPars[i].textContent;
        var locationSplit = locationText.split(', Date: ');
        var location = locationSplit[0].split(': ')[1].trim();
        var date = locationSplit[1].trim();
        calculateLongLat(location, (function (loc, d) {
            return function (latLng) {
                if (map instanceof google.maps.Map) {
                    var marker = new google.maps.Marker({
                        position: latLng,
                        map: map,
                        title: loc + ' - Date: ' + d,
                    });


                    var infoWindow = new google.maps.InfoWindow({
                        content: marker.title,
                    });


                    marker.addListener('click', function () {
                        infoWindow.open(map, marker);
                    });

                    markers.push(marker);
                }
            };
        })(location, date));
    }
    showMap();
}

function calculateLongLat(location, callback) {
    var geocoder = new google.maps.Geocoder();
    var latLng = {lat: 0, lng: 0};
    location = location.replace(/-/g, ' ');
    geocoder.geocode({address: location}, function (results, status) {
        if (status == google.maps.GeocoderStatus.OK) {
            latLng = results[0].geometry.location;
            callback(latLng);
        }
    });
}

function showMap() {
    var mapContainer = document.querySelector("#map");
    mapContainer.style.display = "block";
    setTimeout(function() {
        mapContainer.style.opacity = "1";
        mapContainer.style.visibility = "visible";
        mapContainer.style.transform = "translate(-50%, -50%) scale(1)";
    }, 0);

    var btnContainer = document.querySelector("#close-map-button");
    btnContainer.style.display = "block";
    setTimeout(function() {
        btnContainer.style.opacity = "1";
        btnContainer.style.visibility = "visible";
        btnContainer.style.transform = "translate(-50%, -50%) scale(1)";
    }, 0);
}