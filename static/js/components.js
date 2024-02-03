document.addEventListener('alpine:init', () => {
    Alpine.data("toast", (message) => ({
        visible: false,
        message: "",
        init() {
            this.message = message;
        },
        showToast() {
            setTimeout(() => this.visible = true, 500);
            setTimeout(() => this.visible = false, 3500);
        },
    }));
});