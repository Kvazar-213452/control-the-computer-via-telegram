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
            $(".console").append("good")
        },
        error: function(xhr, status, error) {
            $(".console").append("error dwn folo")
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
        $(".console").append("good")
      },
      error: function(xhr, status, error) {
        $(".console").append("error dwn music")
      }
    });
}
