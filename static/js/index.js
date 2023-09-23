const chat_messages_id = 'chat-messages';

function on_load() {
    const chat_messages_el = document.getElementById(chat_messages_id);
    chat_messages_el.scrollTo({
        top: chat_messages_el.scrollHeight,
        behavior: 'smooth'
    })
}

document.addEventListener('DOMContentLoaded', on_load)