import { RESULT } from "./codes";

function readAsDataURL(file) {
    return new Promise((resolve) => {
        let reader = new FileReader();
        reader.onload = () => {
            resolve(reader.result);
        };
        reader.readAsDataURL(file);
    });
}

function toReadableSize(bytes) {
    if (!bytes) return "0";
    let i = 0;
    while (Math.abs(bytes) >= 1000 && i++ < 2) bytes /= 1000;
    return bytes.toFixed(2) + ["B", "KB", "MB"][Math.min(i, 2)];
}

addEventListener("load", () => {

    let images = {};

    const imageInput = document.getElementById("images"),
        contentInput = document.getElementById("post-content"),
        emailInput = document.getElementById("email"),
        imageSizePreview = document.getElementById("images-size"),
        imagesPreviewBox = document.getElementsByClassName("images-preview-box")[0],
        submitButton = document.getElementById("submit-post"),
        submitResult = document.getElementById("submit-result");

    async function createImagePreviewBox(index) {
        let image = images[index];
        if (!image || !image instanceof Blob) return;
        let previewBox = document.createElement("div"),
            removeButton = document.createElement("button");
        previewBox.classList.add("image-preview-box");
        previewBox.insertAdjacentHTML("afterbegin", `
            <div class="image-preview">
                <img src="${await readAsDataURL(image)}" alt="圖片預覽">
            </div>
            <div class="metadata">
                <p class="name">${image.name}</p>
                <p class="size">${toReadableSize(image.size)}/4MB</p>
            </div>
        `);
        removeButton.classList.add("remove-image", "button-icon");
        removeButton.addEventListener("click", () => {
            delete images[index];
            refreshImagesSize();
            previewBox.parentElement.removeChild(previewBox);
        });
        removeButton.insertAdjacentHTML("afterbegin", `
            <div class="iconify" data-icon="mdi:close"></div>
        `);
        previewBox.appendChild(removeButton);
        return previewBox;
    }

    function getNewIndex() {
        let newIndex = 0;
        for (let index of Object.keys(images)) {
            if (parseInt(index) >= newIndex) newIndex = parseInt(index) + 1;
        }
        return newIndex;
    }

    function getImagesSize() {
        let size = 0;
        for (let index of Object.keys(images)) {
            size += images[index].size;
        }
        return size;
    }

    function refreshImagesSize() {
        let size = toReadableSize(getImagesSize());
        imageSizePreview.innerText = `(${size}/15MB)`;
    }

    imageInput.addEventListener("change", async () => {
        if(imageInput.files?.length) {
            for (let image of imageInput.files) {
                let index = getNewIndex();
                images[index] = image;
                let preview = await createImagePreviewBox(index);
                if (preview) imagesPreviewBox.appendChild(preview);
            }
            refreshImagesSize();
        }
    });

    function createFormData() {
        let form = new FormData();
        form.append("content", contentInput.value || "");
        if (images) {
            for (let index of Object.keys(images)) {
                form.append("images", images[index]);
            }
        }
        if (emailInput) form.append("email", emailInput.value);
        return form;
    }

    function hasInvalid() {
        if (emailInput && !emailInput.value) {
            return "請填寫電子郵件";
        }
        if (!contentInput.value.length && !Object.keys(images).length) {
            return "未填寫任何內容";
        }
        if (getImagesSize() > 15 * (1000 ** 2)) {
            return "所有圖片大小需低於15MB";
        }
        for (let index of Object.keys(images)) {
            if (images[index].size > 4 * (1000 ** 2)) {
                return "單張圖片大小需低於4MB";
            }
        }
        return false;
    }

    function handleResult(result) {
        if (result.code) {
            if (result.code === RESULT.INVALID_QUERY) {
                if (result.details === "content too long") {
                    submitResult.innerText = "貼文內容需少於1000字";
                } else if (result.details === "image too large") {
                    submitResult.innerText = "圖片大小超過限制";
                }
            } else {
                submitResult.innerText = `提交貼文失敗(錯誤代碼：${result.code})`;
            }
        } else {
            location.href = result.redirect_url;
        }
    }

    submitButton.addEventListener("click", async () => {
      let error = hasInvalid();
      if (error) {
          submitResult.innerText = error;
          return;
      }
      let response = await fetch("/post/submit", {
          method: "POST",
          body: createFormData()
      });
      handleResult(await response.json());
    });

});