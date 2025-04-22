type ConnectionStatus = 'connected' | 'connecting' | 'disconnected';
export class SSEClient {
  public status: ConnectionStatus = 'disconnected'
  private eventSource: EventSource | null = null;
  private reconnectInterval: number;
  private maxReconnectAttempts: number;
  private reconnectAttempts = 0;
  
  constructor(
    private url: string,
    options: {
      reconnectInterval?: number;
      maxReconnectAttempts?: number;
    } = {},
  ) {
    this.reconnectInterval = options.reconnectInterval ?? 5000;
    this.maxReconnectAttempts = options.maxReconnectAttempts ?? 5;
  }
  
  connect(): Promise<EventSource> {
    this.status = 'connecting';
    console.log(`Connecting to ${this.url}...`);
    return new Promise((resolve, reject) => {
      this.eventSource = new EventSource(this.url);

      this.eventSource.onopen = () => {
        console.log("Connected!");
        this.status = 'connected';
        this.reconnectAttempts = 0;
        resolve(this.eventSource!);
      };
      
      this.eventSource.onerror = (error) => {
        console.error("SSE connection error", error);
        this.eventSource?.close();
        this.eventSource = null;
        
        if (this.reconnectAttempts < this.maxReconnectAttempts) {
          this.reconnectAttempts++;
          console.log(`Attempting to reconnect (${this.reconnectAttempts}/${this.maxReconnectAttempts})...`);
          setTimeout(this.connect.bind(this), this.reconnectInterval);
        } else {
          console.error("Max reconnect attempts reached");
          reject(new Error("Failed to establish SSE connection with server"));
        }
      };
    })
  }
  
  addEventListener<T>(eventType: string, callback: (data: T) => void): void {
    if (!this.eventSource) {
      throw new Error("SSE client not connected!")
    }
    
    this.eventSource.addEventListener(eventType, ((event: MessageEvent) => {
      try {
        const data = JSON.parse(event.data) as T;
        callback(data);
      } catch (error) {
        console.error("Error parsing SSE event data", error);
        callback(event.data as unknown as T);
      }
    }) as EventListener);
  }
  
  close(): void {
    this.eventSource?.close();
    this.eventSource = null;
    console.log("SSE connection closed");
  }
}
