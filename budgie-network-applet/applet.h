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

#ifndef _ORG_BUDGIE_DESKTOP_NETWORK_APPLET_H_
#define _ORG_BUDGIE_DESKTOP_NETWORK_APPLET_H_

#include <gtk/gtk.h>

void setAllMargins(GtkWidget *widget, gint size);

/* Create a Section Header pseudo Widget
 *
 * @param name      - the name of the section
 * @param icon_name - the name of the icon to use for this section
 *
 * @returns the newly built section header widget
 */
GtkWidget * createSectionHeader(const gchar* name, const gchar* icon_name);

/* Create a Section pseudo Widget
 *
 * @param name      - the name of the section
 * @param icon_name - the name of the icon to use for this section
 * @param list      - the list for this section
 *
 * @returns the newly built section widget
 */
GtkWidget * createSection(const gchar* name, const gchar* icon_name, GtkWidget* list);

GtkWidget * createSettingsWidget();

/* Build the test window and setup the applet inside of it
 *
 * @param app       - the application object for this instance
 * @param user_data - unused
 */
void activate(GtkApplication* app, gpointer user_data);

#endif