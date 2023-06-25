import type { Notification } from './proto/notifications/v1/notifications';

export const deprecated_function_kept_around_incase_I_reintroduce_it = (n: Notification) => {
  const payload = {
    "channel_id": n.channelId,
    "user_id": n.userId,
    "text": n.text
  }

  fetch('http://api.docker.localhost/send-notification', {
    method: 'POST',
    body: JSON.stringify(payload),
    headers: { 'Content-Type': 'application/json' }
  });
};
