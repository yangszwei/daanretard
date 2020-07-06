import { setFormMessage, POST } from "./form";

function getInputSize(input) {
    let size = 0;
    for (let file of input.files || []) {
        size += file.size;
    }
    return size;
}

function readAsDataURL(file) {
    return new Promise((resolve, reject) => {
        let reader = new FileReader();
        reader.onload = () => {
            resolve(reader.result);
        };
        reader.readAsDataURL(file);
    });
}

function initMediaPreview() {
    const media = document.getElementById("media"),
        preview = document.querySelector(".images-preview .container");
    media.addEventListener("change", async (event) => {
        while (preview.firstChild) preview.removeChild(preview.firstChild);
        for (let file of media.files || []) {
            let image = document.createElement("img");
            image.src = await readAsDataURL(file);
            preview.appendChild(image);
        }
    });
}

function initMediaSizePreview() {
    const media = document.getElementById("media"),
        label = document.querySelector(`label[for=media]`);
    media.addEventListener("change", async (event) => {
        let size = getInputSize(media) / (1024 ** 2);
        label.innerText = `上傳圖片 (${size.toPrecision(2)} / 10MB)`;
    });
}

function initSubmitButton() {
    const submit = document.getElementById("submit"),
        email = document.getElementById("email"),
        content = document.getElementById("content"),
        media = document.getElementById("media");
    submit.addEventListener("click", async () => {
        let size = getInputSize(media) / (1024 ** 2);
        if (!email.value) {
            setFormMessage("未登入投稿需驗證電子郵件");
        } else if (!content.value && !media.value) {
            setFormMessage("貼文內容及圖片需至少填寫一項");
        } else if (size > 10) {
            setFormMessage("檔案大小超過限制");
        } else {
            let response = await POST("/post/create", {
                email: email.value || null,
                content: content.value || null,
                media: Array.from(media.files) || []
            });
            let result = await response.json();
            console.log(result)
        }
    });
}

onload = () => {
    initMediaPreview();
    initMediaSizePreview();
    initSubmitButton();
};