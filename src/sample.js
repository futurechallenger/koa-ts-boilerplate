let retWrapper = { ret: "", error: "" };

(function(global) {
  function testHello(e) {
    try {
      return (retWrapper.ret = "method executed returns OK");
    } catch (err) {
      return (retWrapper.error = "Well some went wrong");
    }
  }
})()

