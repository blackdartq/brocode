$("#main-section").click(function(){
    $.getJSON("../getBlogPosts", function(result){
        $.each(result, function(i, field){
              $("section").append(field.Name + "\n" + field.Message + "\n");
        });

    });

});
