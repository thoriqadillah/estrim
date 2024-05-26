import { createSignal, type Component } from 'solid-js';
import { Button } from './ui/button';

const Navbar: Component = () => {
  
  
  return <>
    <div class='w-full px-3 flex items-center justify-between'>
      <div class='font-medium flex gap-2 items-center'>
        <img src="https://api.iconify.design/icon-park-outline:compression.svg?color=%23888888" alt="" />
        <p class='text-sm'>FCompressor</p>
      </div>

      <div>
        <a href="https://github.com">
          <Button variant='link' size='icon' class='flex gap-2 items-center'>
            <img src="https://api.iconify.design/logos:github-icon.svg?color=%23888888" alt="" />
          </Button>
        </a>
      </div>
    </div>
  </>
};

export default Navbar