interface Event {
  subid: string;
  text: string;
  sender: string;
}

export const send_notification = (evt: Event) => {

  fetch('http://api.docker.localhost/send-notification', {
    method: 'POST', 
    body: JSON.stringify(evt), 
    headers: {'Content-Type': 'application/json'}
  })
}
