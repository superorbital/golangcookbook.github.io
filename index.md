---
layout: default
title: Go Cookbook
---

The Go Cookbook is built with the help of people like yourself.  Please support this site and help others by [contributing a recipe of your own](https://github.com/golangcookbook/golangcookbook.github.io#readme).

---

{% for chapter in site.data.chapters %}
### {{ chapter.title }}

  <ul>
  {% for recipe in chapter.recipes %}
    <li>
    {% if recipe.wip %}
      <strong>[work in progress]</strong>
    {% endif %}
    {% if recipe.path %}
      <a href="{{recipe.path}}">{{ recipe.title }}</a>
    {% else %}
      {{ recipe.title }}
    {% endif %}
    </li>
  {% endfor %}
  </ul>
{% endfor %}

