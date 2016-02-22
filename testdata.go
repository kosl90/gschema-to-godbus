package main

const data = `
<?xml version="1.0" encoding="UTF-8"?>
<schemalist>
    <enum id="com.deepin.filemanager.ActivationPolicy">
        <value nick="launch" value="0"/>
        <value nick="display" value="1"/>
        <value nick="ask" value="2"/>
    </enum>

    <enum id="com.deepin.filemanager.ClickPolicy">
        <value nick="single" value="0"/>
        <value nick="double" value="1"/>
    </enum>

    <enum id="com.deepin.filemanager.ViewMode">
        <value nick="icon-view" value="0"/>
        <value nick="list-view" value="1"/>
        <value nick="cascading-view" value="2"/>
    </enum>

    <enum id="com.deepin.filemanager.ShowThumbnailPolicy">
        <value nick="always" value="0"/>
        <value nick="never" value="1"/>
    </enum>

    <enum id="com.deepin.filemanager.PositionPolicy">
        <value nick="bottom" value="0"/>
        <value nick="right" value="1"/>
    </enum>

    <!-- sort and group items like mac -->
    <enum id="com.deepin.filemanager.SortPolicy">
        <value nick="name" value="1"/>
        <value nick="size" value="2"/>
        <value nick="filetype" value="3"/>
        <!-- <value nick="open&#45;with" value="4"/> <!&#45;&#45; the name of open application &#45;&#45;> -->
        <value nick="mtime" value="5"/>
        <!-- <value nick="tag&#45;info" value="6"/> -->
        <!-- <value nick="tag&#45;color" value="7"/> -->
    </enum>

    <enum id="com.deepin.filemanager.SidebarDisplayPolicy">
        <value nick="icon" value="0"/>
        <value nick="icon-and-text" value="1"/>
    </enum>

    <enum id="com.deepin.filemanager.SizeUnit">
        <value nick="Byte" value="0"/>
        <value nick="KiB" value="1"/>
        <value nick="MB" value="2"/>
        <value nick="GB" value="3"/>
        <value nick="TB" value="4"/>
        <value nick="PB" value="5"/>
    </enum>

    <schema id="com.deepin.filemanager" path="/com/deepin/filemanager/">
        <child name="preferences" schema="com.deepin.filemanager.preferences"/>
        <child name="icon-view" schema="com.deepin.filemanager.icon-view"/>
        <child name="list-view" schema="com.deepin.filemanager.list-view"/>
    </schema>

    <schema path="/com/deepin/filemanager/preferences/" id="com.deepin.filemanager.preferences">
        <key name="display-extra-items" type="b">
            <default>false</default>
            <summary>Whether display extra items terminal</summary>
            <description>Whether display extra items in contextmenu.</description>
        </key>
        <key name="activation-policy" enum="com.deepin.filemanager.ActivationPolicy">
            <default>'ask'</default>
            <summary>What to do with executable text files when activated</summary>
            <description>What to do with executable text files when they are activated. Possible values are "launch" to launch them as programs, "ask" to ask what to do via a dialog, and "display" to display them as text files.</description>
        </key>

        <key name="open-dir-in-new-window" type="b">
            <default>false</default>
            <summary>Whether open directory in a new window.</summary>
            <description>Whether open directory in a new window. If not, open directory in a new tab.</description>
        </key>

        <key name="click-policy" enum="com.deepin.filemanager.ClickPolicy">
            <default>'double'</default>
            <summary>Type of click used to launch/open files.</summary>
            <description>Possible values are "single" to launch files on a single click, or "double" to launch theme on a double click.</description>
        </key>

        <key name="view" enum="com.deepin.filemanager.ViewMode">
            <default>'icon-view'</default>
            <summary>The view mode file manager.</summary>
            <description>Possible values are "icon-view", "list-view" and "cascading-view".</description>
        </key>

        <key name="sort-order" enum="com.deepin.filemanager.SortPolicy">
            <default>'name'</default>
            <summary>Default sort policy.</summary>
            <description>Possible values are "name", "size", "filetype", "open-with", "modified_time"("mtime"), "tag-info" and "tag-color"</description>
        </key>

        <key name="show-thumbnail" enum="com.deepin.filemanager.ShowThumbnailPolicy">
            <default>'always'</default>
            <summary>Whether to show thumbnails.</summary>
            <description>Possible values are 'always' to always show thumbnails, 'never' to never show thumbnails.</description>
        </key>

        <key name="thumbnail-size-limitation" type="t">
            <default>52428800</default>
            <summary>The size limitation of thumbnail in bytes.</summary>
            <description>Files over this size won't be thumbnailed. The purpose of this setting is to avoid thumbnailing large files that may take a long time to load or use lots of memory.</description>
        </key>

        <key name="show-zipped-file-in-directory" type="b">
            <default>false</default>
            <summary>Whether to show the content of zipped file in directory.</summary>
            <description>Possible types are zip, tar, tar.gz, bz.</description>
        </key>

        <key name="confirm-trash" type="b">
            <default>false</default>
            <summary>Whether confirm when trash files.</summary>
            <description>If value is true, a confirm dialog will be shown.</description>
        </key>

        <key name="confirm-empty-trash" type="b">
            <default>true</default>
            <summary>Whether confirm when empty trash can.</summary>
            <description>If value is true, a confirm dialog will be shown.</description>
        </key>

        <key name="allow-delete-immediatly" type="b">
            <default>false</default>
            <summary>Whether to enable immediate deletion.</summary>
            <description>Files can be delete immediatly, which means files deleted won't be moved to trash can. This feature can be dangerous, so use caution.</description>
        </key>

        <key name="show-hidden-files" type="b">
            <default>false</default>
            <summary>Whether to show hidden files.</summary>
            <description>If true, hidden files will be shown as well.</description>
        </key>

        <key name="show-hidden-files-in-search-result" type="b">
            <default>false</default>
            <summary>Show hidden file in search result.</summary>
            <description>This option is orthogonal to show-hidden-files.</description>
        </key>

        <key name="show-function-bar" type="b">
            <default>true</default>
            <summary>Whether to show function bar.</summary>
            <description>If true the function bar will be shown if needed.</description>
        </key>

        <key name="label-position" enum="com.deepin.filemanager.PositionPolicy">
            <default>'bottom'</default>
            <summary>The name label position.</summary>
            <description>Possible values are "bottom", "right".</description>
        </key>

        <key name="show-extension-name" type="b">
            <default>true</default>
            <summary>Whether to show extension name</summary>
            <description>If true, show extension name.</description>
        </key>

        <key name="show-brief" type="b">
            <default>false</default>
            <summary>Whether to show brief for files.</summary>
            <description>Brief differs from filetypes, for instances, images will show revolution, files will show size, directory will show item number, symbolic link gets nothing.</description>
        </key>

        <key name="autoplay-audio-video" type="b">
            <default>false</default>
            <summary>Whether to autoplay audio and video.</summary>
            <description>If true, autoplay audio and video when preview.</description>
        </key>

        <key name="sidebar-display-mode" enum="com.deepin.filemanager.SidebarDisplayPolicy">
            <default>'icon-and-text'</default>
            <summary>Display icon or text on sidebar.</summary>
            <description>Possible values are 'icon-and-text' and 'icon'.</description>
        </key>
    </schema>

    <schema path="/com/deepin/filemanager/icon-view/" id="com.deepin.filemanager.icon-view">
        <key name="icon-margin" type="i">
            <range min="40" max="80"/><!-- px -->
            <default>48</default>
            <summary>The margin between two icons.</summary>
            <description>The margin between two icons.</description>
        </key>

        <key name="thumbnail-size" type="i">
            <default>64</default>
            <summary>Default thumbnail icon size.</summary>
            <description>The default size of an icon for a thumbnail in the icon view.</description>
        </key>

        <key name="icon-zoom-level" type="i">
            <range min="50" max="500"/><!-- % -->
            <default>100</default>
            <summary>Default icon zoom level.</summary>
            <description>Defaut icon zoom level used by the icon view.</description>
        </key>
    </schema>

    <flags id="com.deepin.filemanager.list-view-column">
        <value nick="label" value="1"/>
        <value nick="size" value="2"/>
        <value nick="mtime" value="4"/>
        <value nick="filetype" value="8"/>
    </flags>

    <schema path="/com/deepin/filemanager/list-view/" id="com.deepin.filemanager.list-view">
        <key name="thumbnail-size" type="i">
            <default>32</default>
            <summary>Default thumbnail icon size.</summary>
            <description>The default size of an icon for a thumbnail in the list view.</description>
        </key>

        <key name="icon-zoom-level" type="i">
            <range min="50" max="500"/>
            <default>100</default>
            <summary>Default icon zoom level.</summary>
            <description>Defaut icon zoom level used by the list view.</description>
        </key>

        <!-- this maters for non-visible columns -->
        <key name="columns-order" flags="com.deepin.filemanager.list-view-column">
            <default>["label", "size", "mtime", "filetype"]</default>
            <summary>Columns order in list view.</summary>
            <description>Columns order in list view.</description>
        </key>

        <key name="visible-columns" flags="com.deepin.filemanager.list-view-column">
            <default>["label", "size", "mtime", "filetype"]</default>
            <summary>Visible columns in list view.</summary>
            <description>Possible values are "label", "size", "mtime", "filetype".</description>
        </key>
    </schema>
</schemalist>
`
