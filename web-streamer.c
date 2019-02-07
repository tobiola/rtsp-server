#include <gst/gst.h>
#include <gst/rtsp-server/rtsp-server.h>

int main (int argc, char *argv[]) {

  GstRTSPServer *server;
  GMainLoop *loop;

  gst_init(&argc, &argv);

  server = gst_rtsp_server_new();

  /* make a mianloop for the default context */
  loop = g_main_loop_new(NULL, FALSE);

  /* attach the server to the default maincontext */
  gst_rtsp_server_attach(server, NULL);

  g_main_loop_run(loop);

}
