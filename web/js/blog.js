/*$("#main-section").click(function(){
    $.getJSON("../getBlogPosts", function(result){
        $.each(result, function(i, field){
              $("section").append(field.Name + "\n" + field.Message + "\n");
        });

    });

});

function getAllBlogPosts(){
    var requestURL = "http://abyss:8080/getBlogPosts";
    var request = new XMLHttpRequest();
    request.open("GET", requestURL);
    request.responseType = 'json';
    request.send();
    console.log(request);
    throw "testing";

}
*/
