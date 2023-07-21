export interface Message {
  message: string;
  timestamp: string;
  channel: string;
  user: string;
}

export interface OutgoingMessage {
  channelId: string;
  userId: string;
  text: string;
  jwt: string;
}
