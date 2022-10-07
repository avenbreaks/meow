// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import MeowDex from './meow.dex'
import MeowMeow from './meow.meow'


export default { 
  MeowDex: load(MeowDex, 'meow.dex'),
  MeowMeow: load(MeowMeow, 'meow.meow'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}