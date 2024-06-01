import { createSignal, onCleanup, type Component } from 'solid-js';
import { Button } from './components/ui/button';
import Main from './components/main';
import { Icon } from '@iconify-icon/solid';
import { Progress } from './components/ui/progress';

const App: Component = () => {
  const [progress, setProgress] = createSignal(0)
  const timer = setInterval(
    () =>
      setProgress(i => {
        if (i === 100) {
          clearInterval(timer);
          return i;
        }
        return i + 10;
      }),
    400
  );
  onCleanup(() => clearInterval(timer));

  return <>
    <Main>
      <div class='absolute w-full px-4 sm:w-auto top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2'>
        <div class='w-full text-center mb-5'>
          <div class='font-medium flex gap-2 items-center justify-center'>
            <Icon icon='icon-park-outline:compression' class='text-3xl'/>
            <p class='text-3xl font-semibold'>FCompressor</p>
          </div>  
          <p class='text-xs text-center'>
            Background file compressor that lets you compress a file in the background, 
            even without account and the tab is closed, 
            you can still revisit and guaranteed getting the resulted file
          </p>
        </div>

        <Button variant={progress() !== 0 ? 'secondary' : 'default'} disabled={progress() !== 0} class='max-w-lg w-full border relative flex gap-2 mx-auto'>
          <Icon icon='material-symbols:attach-file-add-rounded' class='text-lg' />
          <p>Upload</p>
        </Button>

        <div class='max-w-lg w-full mt-2 mx-auto'>
          <Progress value={progress()}/>
          <div>
            <div class='flex gap-2 justify-between'>
              <p class='text-xs max-w-[10rem] truncate'>Compressing name...</p>
              <p class='text-xs'>{progress()} %</p>
            </div>
          </div>
        </div>
      </div>
    </Main>    
  </>
};

export default App;
