import { reactive } from 'vue';

type Listener<T> = (payload: T) => void
type Events<T = any> = Record<string, Listener<T>[]>

class EventBus {
  private events: Events = reactive({});

  on<T = any>(event: string, listener: Listener<T>) {
    if (!this.events[event]) {
      this.events[event] = [];
    }
    this.events[event].push(listener);
  }

  off<T = any>(event: string, listener: Listener<T>) {
    if (this.events[event]) {
      this.events[event] = this.events[event].filter(l => l !== listener);
    }
  }

  emit<T = any>(event: string, payload?: T) {
    if (this.events[event]) {
      this.events[event].forEach(listener => listener(payload));
    }
  }
}

const eventBus = new EventBus();
export { eventBus }