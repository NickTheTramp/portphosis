/* Set the width of the side navigation to 250px */
function toggleNav() {
    let sidebarWidth = document.getElementById("side-bar").style.width

    if (sidebarWidth === "250px") {
        document.getElementById("side-bar").style.width = "64px";
        document.getElementById("side-bar-toggle-icon").style.transform = "rotate(0deg)";
    } else {
        document.getElementById("side-bar").style.width = "250px";
        document.getElementById("side-bar-toggle-icon").style.transform = "rotate(180deg)";
    }
}