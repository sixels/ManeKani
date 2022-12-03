// import unified from "unified";
import { unified } from 'unified';
import remarkRehype from 'remark-rehype';
import remarkParse from 'remark-parse';
import rehypeStringfy from 'rehype-stringify';

import { HTML } from 'mdast';
import { u } from 'unist-builder';

export default Object.freeze({
  async toHTML(markdown: string): Promise<string> {
    const html = await unified()
      .use(remarkParse)
      .use(remarkRehype, {
        handlers: {
          html(h, node: HTML) {
            const validTags =
              /^<(\/)?(radical|kanji|vocabulary|reading|hiragana|katakana|jpn) *>$/;

            const replaces: { [key: string]: string[] } = {
              radical: ['span', 'class="radical-tag"'],
              kanji: ['span', 'class="kanji-tag"'],
              vocabulary: ['span', 'class="vocabulary-tag"'],
              reading: ['span', 'class="reading-tag"'],

              hiragana: ['span', 'lang="ja"'],
              katakana: ['span', 'lang="ja"'],
              jpn: ['span', 'lang="ja"'],
            };

            const matches = node.value.match(validTags);
            if (!matches) {
              return null;
            }

            const tag: string = matches[2];
            const closing = matches[1] != undefined;

            const replaceWith = closing
              ? replaces[tag][0]
              : replaces[tag].join(' ');

            const newTag = node.value.replace(tag, replaceWith);

            return h.augment(node, u('raw', newTag));
          },
        },
      })
      .use(rehypeStringfy, { allowDangerousHtml: true })
      .process(markdown);

    return html.toString();
  },
});
