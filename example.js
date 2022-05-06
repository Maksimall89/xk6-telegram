import http from "k6/http";
import telegram from "k6/x/telegram";

const conn = telegram.connect(`${__ENV.TOKEN}`, false);
const chatID = 123456789;

export function setup() {
    telegram.send(conn, chatID, "Starting test");
    telegram.sendReplay(conn, chatID, 1, "Replay 1 message from the chat");
}

export default function () {
    http.get('http://test.k6.io');
}

export function teardown() {
    telegram.send(conn, chatID, "Result <b>test</b> with HTML tag");
    telegram.shareImage(conn, chatID, "http://k6.io/logo.png");
    telegram.uploadImagePath(conn, chatID, "/home/k6/logo.png");
    telegram.getUpdate(conn, chatID);
}
