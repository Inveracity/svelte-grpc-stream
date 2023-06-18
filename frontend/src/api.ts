interface Notification {
	subid: string;
	text: string;
	sender: string;
}

export const send_notification = (notif: Notification) => {
	fetch('http://api.docker.localhost/send-notification', {
		method: 'POST',
		body: JSON.stringify(notif),
		headers: { 'Content-Type': 'application/json' }
	});
};
