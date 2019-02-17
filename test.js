import * as streamedian from 'streamedian/player.js';

let mediaElement = rtsp.attach(document.getElementById('test_video'));
let player = new streamedian.WSPlayer(mediaElement, {
    modules: [
        {
            transport: {
               options: {
                   socket: "ws://websocket_proxy_address/ws",
                   errorHandler (e) {
                       alert(`Failed to start player: ${e.message}`);
                   },
                   queryCredentials() {
                       return new Promise((resolve, reject)=>{
                           let c = prompt('input credentials in format user:password');
                           if (c) {
                               this.setCredentials.apply(this, c.split(':'));
                               resolve();
                           } else {
                               reject();
                           }
                       });
                   }
               }
           }
        },
    ]
});
