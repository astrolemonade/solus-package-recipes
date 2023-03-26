/*
 * Copyright 2018 Bryan T. Meyers <bryan@getsol.us>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http: *www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef _ORG_BUDGIE_DESKTOP_NETWORK_APPLET_CONNECTIONS_H_
#define _ORG_BUDGIE_DESKTOP_NETWORK_APPLET_CONNECTIONS_H_

#include <glib-object.h>


enum {
    CONNECTION_STATUS_DISABLED,
    CONNECTION_STATUS_ENABLED
};

enum {
    CONNECTION_NAME,
    CONNECTION_STATUS,
    CONNECTION_STRENGTH,
    CONNECTION_NCOLS
};

G_BEGIN_DECLS

#define NETWORK_TYPE_CONNECTION network_connection_get_type()
G_DECLARE_FINAL_TYPE (NetworkConnection, network_connection, NETWORK, CONNECTION, GObject)

NetworkConnnection *network_connection_new();

G_END_DECLS

#endif