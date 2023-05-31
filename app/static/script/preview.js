let results = document.getElementsByClassName("result-link")
for (const result of results) {
    result.addEventListener("mouseenter", function(e) {
        console.log("mouseenter")
        let el = e.target
        let id = el.id;
        if(id == "result-link"){
            let ifrm = document.getElementById("ifrm")
            ifrm.style.display = "block";
            ifrm.src = el.href;
        }
    })
}
