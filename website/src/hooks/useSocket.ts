export interface WebSocketMessage<T = any> {
  event: '';
  data: T;
}
export type ListenerType = (data: any) => void;
export function useSocket(options: { url: string; onOpen?: Function; onClose?: Function }) {
  const webSocket = new WebSocket(options.url);
  const sendMessage = (event: string, data?: any) => {
    webSocket.send(JSON.stringify({ event, data }));
  };
  const listenersMap = new Map<string, Set<ListenerType>>();
  const on = (event: string, listener: ListenerType) => {
    if (!listenersMap.has(event)) listenersMap.set(event, new Set<(data: any) => void>());
    listenersMap.get(event)?.add(listener);
  };
  const off = (event: string, listener: ListenerType) => {
    listenersMap.get(event)?.delete(listener);
  };
  webSocket.onmessage = (event) => {
    const message = JSON.parse(event.data) as WebSocketMessage;
    listenersMap.get(message.event)?.forEach((it) => {
      it.call(it, message.data);
    });
  };
  webSocket.onopen = () => options.onOpen?.();
  webSocket.onclose = () => options.onClose?.();
  return {
    on,
    off,
    sendMessage,
  };
}
