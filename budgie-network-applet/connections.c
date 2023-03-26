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

#include "connections.h"

typedef struct {
    gchar *name;
    gint  status;
    gint  strength;
} NetworkConnectionPrivate;

struct _NetworkConnection {
    GObject parent_instance;

    NetworkConnectionPrivate *priv;
}

G_DEFINE_TYPE_WITH_PRIVATE(NetworkConnection, network_connection, G_TYPE_OBJECT);

static void network_connection_class_init(NetworkConnectionClass *klass) {}

static void network_connection_init(NetworkConnection *self) {
    NetworkConnectionPrivate *priv = network_connection_get_instance_private(self);
    priv->name     = "";
    priv->status   = CONNECTION_STATUS_DISABLED;
    priv->strength = 0;
}

GObject *network_connection_new(gchar){
    GObject self = g_object_new(NETWORK_TYPE_CONNECTION);
}

#endif