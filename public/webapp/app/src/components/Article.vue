<template>
  <v-container>
    <v-dialog
    v-model="dialog"
    width="500"
    >
        <v-card>
            <v-card-title
            class="headline grey lighten-2"
            primary-title
            >
            Add new section
            </v-card-title>

            <v-card-text>
                <v-select
                :items="contentTypes"
                label="Content Type"
                v-model="choosenType"
                required
                ></v-select>
            </v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
                color="primary"
                flat
                @click="dismissSection()"
            >
                Cancel
            </v-btn>
            <v-btn
                color="primary"
                flat
                :disabled="!choosenType"
                @click="addSection()"
            >
                Add
            </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
    <v-layout
      text-xs-center
      wrap
    >
      <v-flex xs12>
          <h1 v-if="$route.params.id">Edit article</h1>
          <h1 v-else>New article</h1>
      </v-flex>

      <v-flex xs12>
          <v-form>
            <v-container>
                <v-layout>
                    <v-flex
                    xs12
                    md12
                    >
                    <v-text-field
                        v-model="model.title"
                        label="Title"
                        required
                    ></v-text-field>
                    </v-flex>
                </v-layout>
                <v-layout v-for="(selection, index) in model.sections"
                    :key="index">
                        <v-textarea
                            v-model="selection.content"
                        ></v-textarea>
                </v-layout>
                <v-layout>
                    <v-flex
                        xs12
                        md12
                        >
                        <v-btn 
                        flat 
                        icon
                        large
                        v-on:click="showDialog"
                        color="info">
                            +
                        </v-btn>
                    </v-flex>
                </v-layout>
                <v-layout>
                    <v-btn
                        color="primary"
                        flat
                        @click="save()"
                    >
                        Save
                    </v-btn>
                </v-layout>
            </v-container>
        </v-form>
      </v-flex>
    


    </v-layout>
  </v-container>
</template>

<script>
  export default {
    data: () => ({
        contentTypes: ['image', 'source-code', 'paragraph'],
        choosenType: '',
        dialog: false,
        model: {  
            title: '',
            sections: []
        }
    }),
    methods: {
        showDialog: function () {
            this.dialog = true;
        },
        addSection: function () {
            if (this.choosenType) {
                this.model.sections.push({ type: this.choosenType, content: '' });
                this.choosenType = '';
                this.dialog = false;
            }
        },
        dismissSection: function () {
            this.choosenType = '';
            this.dialog = false;
        },
        save: function () {
            console.log('Saving...');
        }
    }
  }
</script>

<style>

</style>
