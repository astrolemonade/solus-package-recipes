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

#include "connection_list.h"

static GtkWidget * vpnConnectionItem(gchar *name, int connected) {
    GtkWidget * box;
    GtkWidget * label;
    GtkWidget * button = NULL;
    GtkWidget * icon;
    GtkStyleContext *style;

    label = gtk_label_new(name);
    gtk_widget_set_hexpand(label, TRUE);
    gtk_widget_set_halign(label, GTK_ALIGN_START);

    if(connected==1){
        button = gtk_button_new_from_icon_name("cancel",GTK_ICON_SIZE_BUTTON);
        style = gtk_widget_get_style_context(button);
        gtk_style_context_add_class(style, "flat");
        gtk_style_context_remove_class(style, "image_button");
    }

    if(connected == 1) {
        icon = gtk_image_new_from_icon_name("network-vpn-symbolic",GTK_ICON_SIZE_BUTTON);
    } else {
        icon = gtk_image_new_from_icon_name("network-vpn-offline-symbolic",GTK_ICON_SIZE_BUTTON);
    }

    box = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 0);
    gtk_box_pack_start(GTK_BOX(box), label, TRUE, TRUE, 5);
    if(button)
        gtk_box_pack_start(GTK_BOX(box), button, FALSE, FALSE, 5);
    gtk_box_pack_start(GTK_BOX(box), icon, FALSE, FALSE, 5);

    return box;
}

static GtkWidget * cellularConnectionItem(gchar *name, int connected, int signal) {
    GtkWidget * box;
    GtkWidget * label;
    GtkWidget * button = NULL;
    GtkWidget * icon;
    GtkStyleContext *style;

    label = gtk_label_new(name);
    gtk_widget_set_hexpand(label, TRUE);
    gtk_widget_set_halign(label, GTK_ALIGN_START);

    if(connected==1){
        button = gtk_button_new_from_icon_name("cancel",GTK_ICON_SIZE_BUTTON);
        style = gtk_widget_get_style_context(button);
        gtk_style_context_add_class(style, "flat");
        gtk_style_context_remove_class(style, "image_button");
    }
    if(signal > 80) {
        icon = gtk_image_new_from_icon_name("network-cellular-signal-excellent",GTK_ICON_SIZE_BUTTON);
    } else if(signal > 60) {
        icon = gtk_image_new_from_icon_name("network-cellular-signal-good",GTK_ICON_SIZE_BUTTON);
    } else if(signal > 40) {
        icon = gtk_image_new_from_icon_name("network-cellular-signal-low",GTK_ICON_SIZE_BUTTON);
    } else {
        icon = gtk_image_new_from_icon_name("network-cellular-signal-none",GTK_ICON_SIZE_BUTTON);
    }

    box = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 0);
    gtk_box_pack_start(GTK_BOX(box), label, TRUE, TRUE, 5);
    if(button)
        gtk_box_pack_start(GTK_BOX(box), button, FALSE, FALSE, 5);
    gtk_box_pack_start(GTK_BOX(box), icon, FALSE, FALSE, 5);

    return box;
}

static GtkWidget * wifiConnectionItem(gchar *name, int connected, int signal) {
    GtkWidget * box;
    GtkWidget * label;
    GtkWidget * button = NULL;
    GtkWidget * icon;
    GtkStyleContext *style;

    label = gtk_label_new(name);
    gtk_widget_set_hexpand(label, TRUE);
    gtk_widget_set_halign(label, GTK_ALIGN_START);

    if(connected==1){
        button = gtk_button_new_from_icon_name("cancel",GTK_ICON_SIZE_BUTTON);
        style = gtk_widget_get_style_context(button);
        gtk_style_context_add_class(style, "flat");
        gtk_style_context_remove_class(style, "image_button");
    }
    if(signal > 80) {
        icon = gtk_image_new_from_icon_name("network-wireless-signal-excellent",GTK_ICON_SIZE_BUTTON);
    } else if(signal > 60) {
        icon = gtk_image_new_from_icon_name("network-wireless-signal-good",GTK_ICON_SIZE_BUTTON);
    } else if(signal > 40) {
        icon = gtk_image_new_from_icon_name("network-wireless-signal-low",GTK_ICON_SIZE_BUTTON);
    } else {
        icon = gtk_image_new_from_icon_name("network-wireless-signal-none",GTK_ICON_SIZE_BUTTON);
    }

    box = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 0);
    gtk_box_pack_start(GTK_BOX(box), label, TRUE, TRUE, 5);
    if(button)
        gtk_box_pack_start(GTK_BOX(box), button, FALSE, FALSE, 5);
    gtk_box_pack_start(GTK_BOX(box), icon, FALSE, FALSE, 5);

    return box;
}

static GtkWidget * wiredConnectionItem(gchar *name, int enabled, int connected) {
    GtkWidget * box;
    GtkWidget * label;
    GtkWidget * button = NULL;
    GtkWidget * icon;
    GtkStyleContext *style;

    label = gtk_label_new(name);
    gtk_widget_set_hexpand(label, TRUE);
    gtk_widget_set_halign(label, GTK_ALIGN_START);

    if((enabled == 1) && (connected==1)){
        button = gtk_button_new_from_icon_name("reload",GTK_ICON_SIZE_BUTTON);
        style = gtk_widget_get_style_context(button);
        gtk_style_context_add_class(style, "flat");
        gtk_style_context_remove_class(style, "image_button");
    }

    if(connected == 1) {
        icon = gtk_image_new_from_icon_name("network-transmit-symbolic",GTK_ICON_SIZE_BUTTON);
    } else {
        icon = gtk_image_new_from_icon_name("network-idle-symbolic",GTK_ICON_SIZE_BUTTON);
    }

    box = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 0);
    gtk_box_pack_start(GTK_BOX(box), label, TRUE, TRUE, 5);
    if(button)
        gtk_box_pack_start(GTK_BOX(box), button, FALSE, FALSE, 5);
    gtk_box_pack_start(GTK_BOX(box), icon, FALSE, FALSE, 5);

    return box;
}

GtkWidget * fakeVPNConnectionList(){
    GtkWidget *list;
    GtkWidget *item1;
    GtkWidget *item2;

    item1 = vpnConnectionItem("RIT", 1);
    item2 = vpnConnectionItem("Private Internet Access", 0);

    list = gtk_list_box_new();
    gtk_list_box_insert(GTK_LIST_BOX(list), item1, -1);
    gtk_list_box_insert(GTK_LIST_BOX(list), item2, -1);
    gtk_list_box_set_selection_mode(GTK_LIST_BOX(list), GTK_SELECTION_NONE);
    return list;
}

GtkWidget * fakeCellularConnectionList(){
    GtkWidget *list;
    GtkWidget *item1;
    GtkWidget *item2;

    item1 = cellularConnectionItem("Verizon", 1, 95);
    item2 = cellularConnectionItem("Sprint", 0, 60);

    list = gtk_list_box_new();
    gtk_list_box_insert(GTK_LIST_BOX(list), item1, -1);
    gtk_list_box_insert(GTK_LIST_BOX(list), item2, -1);
    gtk_list_box_set_selection_mode(GTK_LIST_BOX(list), GTK_SELECTION_NONE);
    return list;
}

GtkWidget * fakeWifiConnectionList(){
    GtkWidget *list;
    GtkWidget *item1;
    GtkWidget *item2;

    item1 = wifiConnectionItem("Connected AP, Med", 1, 50);
    item2 = wifiConnectionItem("Low AP", 0, 20);

    list = gtk_list_box_new();
    gtk_list_box_insert(GTK_LIST_BOX(list), item1, -1);
    gtk_list_box_insert(GTK_LIST_BOX(list), item2, -1);
    gtk_list_box_set_selection_mode(GTK_LIST_BOX(list), GTK_SELECTION_NONE);
    return list;
}

GtkWidget * fakeWiredConnectionList(){
    GtkWidget *list;
    GtkWidget *item1;
    GtkWidget *item2;

    item1 = wiredConnectionItem("eth0", 1, 1);
    item2 = wiredConnectionItem("eth1", 1, 0);

    list = gtk_list_box_new();
    gtk_list_box_insert(GTK_LIST_BOX(list), item1, -1);
    gtk_list_box_insert(GTK_LIST_BOX(list), item2, -1);
    gtk_list_box_set_selection_mode(GTK_LIST_BOX(list), GTK_SELECTION_NONE);
    return list;
}

