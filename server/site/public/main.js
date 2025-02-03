function upload_foto() {
    let formData = new FormData();
    formData.append('text', $('#text_foto').val());
    formData.append('file', $('#file_foto')[0].files[0]);

    $.ajax({
        url: '/php/upload_foto.php',
        type: 'POST',
        data: formData,
        contentType: false,
        processData: false,
        success: function(response) {
            $(".console").append("<p>-- good</p>");
            get_data();
        },
        error: function(xhr, status, error) {
            $(".console").append("<p>-- error dwn folo</p>");
        }
    });
}

function upload_music() {
    let formData = new FormData();
    formData.append('text', $('#text_music').val());
    formData.append('file', $('#file_music')[0].files[0]);

    $.ajax({
      url: '/php/upload_music.php',
      type: 'POST',
      data: formData,
      contentType: false,
      processData: false,
      success: function(response) {
        $(".console").append("<p>-- good</p>");
        get_data();
      },
      error: function(xhr, status, error) {
        $(".console").append("<p>-- error dwn music</p>");
      }
    });
}

function get_data() {
    $(".console").append("<p>-- get_data</p>")

    $.get('/php/music.php', function(data) {
        $("#data_music").html(null)

        for (let key in data) {
            if (data.hasOwnProperty(key)) {
                let text = `<div class="item"><p onclick="del_music('${key}')">del</p> ${key}: <span>${data[key]}</span></div>`;

                $("#data_music").append(text)
            }
        }
    });

    $.get('/php/foto.php', function(data) {
        $("#data_foto").html(null)

        for (let key in data) {
            if (data.hasOwnProperty(key)) {
                let text = `<div class="item"><p onclick="del_foto('${key}')">del</p> ${key}: <span>${data[key]}</span></div>`;

                $("#data_foto").append(text)
            }
        }
    });
}

function del_music(text_) {
    console.log(text_)
    let data = {text: text_};

    $.ajax({
        url: "/php/del_music.php",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            $(".console").append("<p>-- good del</p>")
            get_data();
        }
    });
}

function del_foto(text_) {
    console.log(text_)
    let data = {text: text_};

    $.ajax({
        url: "/php/del_foto.php",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
            $(".console").append("<p>-- good del</p>")
            get_data();
        }
    });
}
