
const API_BASE = window.location.origin;
const REGISTER_ENDPOINT = `${API_BASE}/api/v1/url/register`;

const urlInput = document.getElementById('urlInput');
const shortenBtn = document.getElementById('shortenBtn');
const errorMsg = document.getElementById('errorMsg');
const resultPanel = document.getElementById('resultPanel');
const shortLink = document.getElementById('shortLink');
const targetPath = document.getElementById('targetPath');
const copyBtn = document.getElementById('copyBtn');

urlInput.addEventListener('keydown', (e) => {
    if (e.key === 'Enter') shortenURL();
});

function isValidURL(str) {
    try {
        const url = new URL(str);
        return url.protocol === 'http:' || url.protocol === 'https:';
    } catch {
        return false;
    }
}

function showError(msg) {
    errorMsg.textContent = msg;
    errorMsg.classList.add('visible');
}

function clearError() {
    errorMsg.textContent = '';
    errorMsg.classList.remove('visible');
}

async function shortenURL() {
    clearError();
    const raw = urlInput.value.trim();

    if (!raw) {
        showError('Please enter a URL.');
        return;
    }

    if (!isValidURL(raw)) {
        showError('Enter a valid URL starting with http')
        return;
    }


    shortenBtn.disabled = true;
    shortenBtn.innerHTML = '<span class="spinner"></span>';

    try {
        const res = await fetch(REGISTER_ENDPOINT, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ target_path: raw }),
        });

        if (res.status === 429) {
            showError('Rate limited — please wait a moment and try again.');
            return;
        }

        if (!res.ok) {
            const body = await res.json().catch(() => null);
            showError(body?.error || `Server returned ${res.status}`);
            return;
        }

        const data = await res.json();


        shortLink.href = data.short_path;
        shortLink.textContent = data.short_path;
        targetPath.textContent = data.target_path;
        resultPanel.classList.add('visible');


        copyBtn.textContent = 'Copy';
        copyBtn.classList.remove('copied');

    } catch (err) {
        showError('Could not reach the server. Is it running?');
        console.error(err);
    } finally {
        shortenBtn.disabled = false;
        shortenBtn.textContent = 'Shorten';
    }
}

async function copyURL() {
    const url = shortLink.textContent;
    try {
        await navigator.clipboard.writeText(url);
        copyBtn.textContent = 'Copied!';
        copyBtn.classList.add('copied');
        setTimeout(() => {
            copyBtn.textContent = 'Copy';
            copyBtn.classList.remove('copied');
        }, 2000);
    } catch {

        const ta = document.createElement('textarea');
        ta.value = url;
        document.body.appendChild(ta);
        ta.select();
        document.execCommand('copy');
        document.body.removeChild(ta);
        copyBtn.textContent = 'Copied!';
        copyBtn.classList.add('copied');
    }
}
