import { Notification } from "./proto/notifications/v1/notifications";

interface Event {
  subid: string;
  text: string;
  sender: string;
}

export const send_notification = (evt: Notification) => {

  fetch('http://api.docker.localhost/send-notification', {
    method: 'POST',
    body: JSON.stringify(evt),
    headers: {'Content-Type': 'application/json'}
  })
}
