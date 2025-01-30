function upload_foto() {
    var formData = new FormData();
    formData.append('text', $('#text_foto').val());
    formData.append('file', $('#file_foto')[0].files[0]);

    $.ajax({
        url: '/upload_foto',
        type: 'POST',
        data: formData,
        contentType: false,
        processData: false,
        success: function(response) {
            $(".console").append("<p>good</p>")
        },
        error: function(xhr, status, error) {
            $(".console").append("<p>error dwn folo</p>")
        }
    });
}

function upload_music() {
    var formData = new FormData();
    formData.append('text', $('#text_music').val());
    formData.append('file', $('#file_music')[0].files[0]);

    $.ajax({
      url: '/upload_music',
      type: 'POST',
      data: formData,
      contentType: false,
      processData: false,
      success: function(response) {
        $(".console").append("<p>good</p>")
      },
      error: function(xhr, status, error) {
        $(".console").append("<p>error dwn music</p>")
      }
    });
}

function get_data() {
    $(".console").append("<p>get_data</p>")

    $.get('/music', function(data) {
        $("#data_music").html(null)

        for (let key in data) {
            if (data.hasOwnProperty(key)) {
                let text = `<div class="item"><p>del</p> ${key}: <span>${data[key]}</span></div>`;

                $("#data_music").append(text)
            }
        }
    });

    $.get('/foto', function(data) {
        $("#data_foto").html(null)

        for (let key in data) {
            if (data.hasOwnProperty(key)) {
                let text = `<div class="item"><p>del</p> ${key}: <span>${data[key]}</span></div>`;

                $("#data_foto").append(text)
            }
        }
    });
}
