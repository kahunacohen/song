document.addEventListener('alpine:init', () => {
    Alpine.store("flash", {
        msg: "",
        on: false
})
Alpine.data("flash", (msg) => ({		
    msg,		
    init() { 
        this.$watch("$store.flash.on", (on) => {
            if (on) {
                setTimeout(() => {
                    this.$store.flash.on = false;
                    this.$store.flash.msg = "";
                }, 3500);
            }
        })
    }
}))
});