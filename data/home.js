
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

// Save input text when it's refresh
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
