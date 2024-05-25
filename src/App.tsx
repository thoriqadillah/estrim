import { createSignal, type Component } from 'solid-js';
import { Button } from './components/ui/button';

const App: Component = () => {
  const [count, setCount] = createSignal(0)
  return <>
    <div class='flex justify-center h-screen items-center'>
      <Button onClick={() => setCount(count() + 1)}>Counter {count()}</Button>
    </div>
  </>
};

export default App;
