let pass = document.querySelector("#pass")
    let togglePass = document.querySelector("#togglePass")
    togglePass.addEventListener("click", (e) => {
    if (pass.type === "password") {
    pass.type = "text"
    togglePass.textContent = "Hide"
} else {
    pass.type = "password"
    togglePass.textContent = "Show"
}
})
