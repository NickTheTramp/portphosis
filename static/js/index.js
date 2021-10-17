/* Set the width of the side navigation to 250px */
function toggleNav() {
    let sidebarWidth = document.getElementById("side-bar").style.width

    if (sidebarWidth === "400px") {
        document.getElementById("side-bar").style.width = "64px";
        document.getElementById("side-bar-toggle-icon").style.transform = "rotate(0deg)";
    } else {
        document.getElementById("side-bar").style.width = "400px";
        document.getElementById("side-bar-toggle-icon").style.transform = "rotate(180deg)";
    }
}

function toggleContainer(configurationName) {
    let url = '/toggle?name=' + configurationName;

    fetch(url)
        .then(response => response)
        .then(data => {
            console.log('Success:', data);
        })
        .catch((error) => {
            console.error('Error:', error);
        });
}

let menuItems = document.getElementsByClassName("side-bar-menu-item-icon")

for (let menuItem of menuItems) {
    menuItem.addEventListener("click", toggleContainer);
}
