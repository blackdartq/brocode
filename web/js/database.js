function getAllBlogPosts(){
    var requestURL = "/getBlogPosts";
    var request = new XMLHttpRequest();
    request.open("GET", requestURL, false);
    request.send();
    return JSON.parse(request.responseText);
}

window.onload = function(){
    var app = new Vue({
        el: `#app`,
        data:{
            message: getAllBlogPosts()[0].Name
        }
    });
}

