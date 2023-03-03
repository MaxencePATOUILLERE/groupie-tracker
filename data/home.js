const cards = document.querySelectorAll(".card__inner");

for (let i = 0; i < cards.length; i++) {
    const card = cards[i];
    card.addEventListener("click", () =>
    {
        card.classList.toggle("is-flipped");

    });
}

function popup() {
    document.getElementsByClassName("locationDate").display = "block";
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

        if (filter === "") {
            li[i].style.display = "none";
        } else if (txtValue.toUpperCase().indexOf(filter) > -1) {
            console.log("test")
            li[i].style.display = "block";
        }else {
            li[i].style.display = "none";
            }
    }
}