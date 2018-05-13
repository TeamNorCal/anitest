$(document).ready(function() {
  console.log("Calling init");
  $.ajax("/init")
  // .then(function(data, status) {
  //   console.log("init ran. Status: " + status);
  //   return $.getJSON("/getFrame")
  // })
  // .done(function(data, status) {
    // console.log("Got data", frame);
    // anim.processFrame(frame);
  // })
  .done(function(data, status) {
    setInterval(function() {
      //console.log("Getting frame");
      $.getJSON("/getFrame", function(frame) {
        //console.log("Processing frame data", frame.Data[0]);
        anim.processFrame(frame);
      });
    }, 30);
  })
  .fail(function(jqXHR, status, errorThrown) {
    console.error("Failed with status " + status + " and error " + errorThrown);
  });
});

var anim = {
  processFrame: function(frame) {
    var data = frame.Data;
    var pixels = $(".pixel");
    data.forEach(function(pd, idx) {
      var color = "rgba(" + pd.R + ", " + pd.G + ", " + pd.B + ", " + pd.A + ")";
      pixels[idx].style["backgroundColor"] = color;
    });
  }
};
